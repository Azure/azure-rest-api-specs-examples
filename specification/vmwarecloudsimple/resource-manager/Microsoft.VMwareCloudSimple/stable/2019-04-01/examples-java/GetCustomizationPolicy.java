/** Samples for CustomizationPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetCustomizationPolicy.json
     */
    /**
     * Sample code: GetCustomizationPolicy.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getCustomizationPolicy(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .customizationPolicies()
            .getWithResponse("myResourceGroup", "myPrivateCloud", "Linux1", com.azure.core.util.Context.NONE);
    }
}
