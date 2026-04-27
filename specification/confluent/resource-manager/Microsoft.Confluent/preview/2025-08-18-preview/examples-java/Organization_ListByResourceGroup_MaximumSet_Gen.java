
/**
 * Samples for Organization ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListByResourceGroupMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listByResourceGroup("rgconfluent", com.azure.core.util.Context.NONE);
    }
}
