
/**
 * Samples for Organization Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/Organization_Delete_MinimumSet_Gen.json
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
