/** Samples for HanaInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_Get.json
     */
    /**
     * Sample code: Get properties of a HANA instance.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void getPropertiesOfAHANAInstance(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager
            .hanaInstances()
            .getByResourceGroupWithResponse("myResourceGroup", "myHanaInstance", com.azure.core.util.Context.NONE);
    }
}
