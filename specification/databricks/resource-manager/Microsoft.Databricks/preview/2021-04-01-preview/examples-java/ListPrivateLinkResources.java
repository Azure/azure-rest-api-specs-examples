/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/ListPrivateLinkResources.json
     */
    /**
     * Sample code: List private link resources.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listPrivateLinkResources(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.privateLinkResources().list("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
