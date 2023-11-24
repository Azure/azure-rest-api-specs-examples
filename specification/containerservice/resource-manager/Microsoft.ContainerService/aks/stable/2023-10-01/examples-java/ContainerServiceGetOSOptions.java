/** Samples for ManagedClusters GetOSOptions. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/ContainerServiceGetOSOptions.json
     */
    /**
     * Sample code: Get Container Service OS Options.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getContainerServiceOSOptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .getOSOptionsWithResponse("location1", null, com.azure.core.util.Context.NONE);
    }
}
