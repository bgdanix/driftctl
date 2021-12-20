package google

import "github.com/snyk/driftctl/pkg/resource"

const GoogleComputeNetworkResourceType = "google_compute_network"

func initGoogleComputeNetworkMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(GoogleComputeNetworkResourceType, func(res *resource.Resource) {
		res.Attributes().SafeDelete([]string{"timeouts"})
		res.Attributes().SafeDelete([]string{"self_link"})
		res.Attributes().SafeDelete([]string{"gateway_ipv4"})
		res.Attributes().SafeDelete([]string{"delete_default_routes_on_create"})
	})
	resourceSchemaRepository.SetResolveReadAttributesFunc(GoogleComputeNetworkResourceType, func(res *resource.Resource) map[string]string {
		return map[string]string{
			"name": *res.Attributes().GetString("name"),
		}
	})
	resourceSchemaRepository.SetFlags(GoogleComputeNetworkResourceType, resource.FlagDeepMode)
}
