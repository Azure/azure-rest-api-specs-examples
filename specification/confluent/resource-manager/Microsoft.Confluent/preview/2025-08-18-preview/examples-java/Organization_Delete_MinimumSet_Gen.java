
/**
 * Samples for Organization Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_Delete_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationDeleteMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().delete("rgconfluent", "w", com.azure.core.util.Context.NONE);
    }
}
