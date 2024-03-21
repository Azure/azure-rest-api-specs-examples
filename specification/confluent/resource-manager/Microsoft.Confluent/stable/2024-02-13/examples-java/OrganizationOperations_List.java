
/**
 * Samples for OrganizationOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * OrganizationOperations_List.json
     */
    /**
     * Sample code: OrganizationOperations_List.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationOperationsList(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizationOperations().list(com.azure.core.util.Context.NONE);
    }
}
