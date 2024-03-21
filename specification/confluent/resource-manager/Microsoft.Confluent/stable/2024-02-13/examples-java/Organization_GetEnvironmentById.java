
/**
 * Samples for Organization GetEnvironmentById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_GetEnvironmentById.json
     */
    /**
     * Sample code: Organization_GetEnvironmentById.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetEnvironmentById(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getEnvironmentByIdWithResponse("myResourceGroup", "myOrganization", "dlz-f3a90de",
            com.azure.core.util.Context.NONE);
    }
}
