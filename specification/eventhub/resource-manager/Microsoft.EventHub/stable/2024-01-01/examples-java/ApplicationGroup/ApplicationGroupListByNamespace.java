
/**
 * Samples for ApplicationGroup ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/
     * ApplicationGroupListByNamespace.json
     */
    /**
     * Sample code: ListApplicationGroups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listApplicationGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getApplicationGroups().listByNamespace("contosotest",
            "contoso-ua-test-eh-system-1", com.azure.core.util.Context.NONE);
    }
}
