
/**
 * Samples for Organization ListEnvironments.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_EnvironmentList.json
     */
    /**
     * Sample code: Organization_ListEnvironments.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListEnvironments(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listEnvironments("myResourceGroup", "myOrganization", 10, null,
            com.azure.core.util.Context.NONE);
    }
}
