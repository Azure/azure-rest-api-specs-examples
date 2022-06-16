import com.azure.core.util.Context;

/** Samples for Organization GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Get.json
     */
    /**
     * Sample code: Organization_Get.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationGet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getByResourceGroupWithResponse("myResourceGroup", "myOrganization", Context.NONE);
    }
}
