
/**
 * Samples for Apms List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apms_List.json
     */
    /**
     * Sample code: Apms_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apmsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApms().list("myResourceGroup", "myservice",
            com.azure.core.util.Context.NONE);
    }
}
