
/**
 * Samples for Organization List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_ListBySubscription.json
     */
    /**
     * Sample code: Organization_ListBySubscription.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListBySubscription(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
