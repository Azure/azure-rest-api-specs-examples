/** Samples for HanaInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_Start.json
     */
    /**
     * Sample code: Start a HANA instance.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void startAHANAInstance(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.hanaInstances().start("myResourceGroup", "myHanaInstance", com.azure.core.util.Context.NONE);
    }
}
