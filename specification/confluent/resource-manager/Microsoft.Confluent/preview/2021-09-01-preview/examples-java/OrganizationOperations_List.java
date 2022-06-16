import com.azure.core.util.Context;

/** Samples for OrganizationOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/OrganizationOperations_List.json
     */
    /**
     * Sample code: OrganizationOperations_List.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationOperationsList(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizationOperations().list(Context.NONE);
    }
}
