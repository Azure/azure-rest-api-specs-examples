
/**
 * Samples for Organization GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_Get.json
     */
    /**
     * Sample code: Organization_Get.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getByResourceGroupWithResponse("myResourceGroup", "myOrganization",
            com.azure.core.util.Context.NONE);
    }
}
