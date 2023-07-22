/** Samples for Catalogs ListDeployments. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeploymentsByCatalog.json
     */
    /**
     * Sample code: Catalogs_ListDeployments.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsListDeployments(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .catalogs()
            .listDeployments(
                "MyResourceGroup1", "MyCatalog1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
