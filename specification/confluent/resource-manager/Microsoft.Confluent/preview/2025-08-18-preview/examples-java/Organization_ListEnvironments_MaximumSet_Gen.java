
/**
 * Samples for Organization ListEnvironments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_ListEnvironments_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListEnvironments_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListEnvironmentsMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listEnvironments("rgconfluent", "zgvcszgobzkrvomvhkabzamqincp", 21, "e",
            com.azure.core.util.Context.NONE);
    }
}
