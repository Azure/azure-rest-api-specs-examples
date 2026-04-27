
/**
 * Samples for Organization GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_Get_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgconfluent", "nnyqgkogkmwjubhfaynme",
            com.azure.core.util.Context.NONE);
    }
}
