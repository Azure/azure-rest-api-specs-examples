
/**
 * Samples for OrganizationOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/OrganizationOperations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: OrganizationOperations_List_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationOperationsListMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizationOperations().list(com.azure.core.util.Context.NONE);
    }
}
