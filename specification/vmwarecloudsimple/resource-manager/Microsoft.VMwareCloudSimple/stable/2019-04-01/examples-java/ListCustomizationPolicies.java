/** Samples for CustomizationPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListCustomizationPolicies.json
     */
    /**
     * Sample code: ListCustomizationPolicies.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listCustomizationPolicies(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .customizationPolicies()
            .list("myResourceGroup", "myPrivateCloud", null, com.azure.core.util.Context.NONE);
    }
}
