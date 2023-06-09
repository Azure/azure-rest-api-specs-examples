/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Get a private link resource.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void getAPrivateLinkResource(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .privateLinkResources()
            .getWithResponse("myResourceGroup", "myWorkspace", "databricks_ui_api", com.azure.core.util.Context.NONE);
    }
}
