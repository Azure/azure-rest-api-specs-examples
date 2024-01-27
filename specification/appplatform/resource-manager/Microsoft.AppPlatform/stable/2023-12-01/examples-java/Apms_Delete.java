
/**
 * Samples for Apms Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apms_Delete.json
     */
    /**
     * Sample code: Apms_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apmsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApms().delete("myResourceGroup", "myservice",
            "myappinsights", com.azure.core.util.Context.NONE);
    }
}
