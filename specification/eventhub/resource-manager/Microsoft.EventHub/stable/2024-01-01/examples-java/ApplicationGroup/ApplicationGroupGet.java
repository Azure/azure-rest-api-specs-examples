
/**
 * Samples for ApplicationGroup Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/
     * ApplicationGroupGet.json
     */
    /**
     * Sample code: ApplicationGroupGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationGroupGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getApplicationGroups().getWithResponse("contosotest",
            "contoso-ua-test-eh-system-1", "appGroup1", com.azure.core.util.Context.NONE);
    }
}
