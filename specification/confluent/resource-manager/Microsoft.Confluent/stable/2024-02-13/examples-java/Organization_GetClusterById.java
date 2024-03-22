
/**
 * Samples for Organization GetClusterById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * Organization_GetClusterById.json
     */
    /**
     * Sample code: Organization_GetClusterById.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGetClusterById(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getClusterByIdWithResponse("myResourceGroup", "myOrganization", "env-12132",
            "dlz-f3a90de", com.azure.core.util.Context.NONE);
    }
}
