
/**
 * Samples for Organization ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_ListByResourceGroup.json
     */
    /**
     * Sample code: Organization_ListByResourceGroup.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListByResourceGroup(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
