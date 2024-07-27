
/**
 * Samples for ApplicationGroup Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/
     * ApplicationGroupDelete.json
     */
    /**
     * Sample code: ApplicationGroupDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationGroupDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getApplicationGroups().deleteWithResponse("contosotest",
            "contoso-ua-test-eh-system-1", "appGroup1", com.azure.core.util.Context.NONE);
    }
}
