
/**
 * Samples for OrganizationOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/OrganizationOperations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: OrganizationOperations_List_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationOperationsListMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizationOperations().list(com.azure.core.util.Context.NONE);
    }
}
