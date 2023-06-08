/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/ListPrivateLinkResources.json
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
