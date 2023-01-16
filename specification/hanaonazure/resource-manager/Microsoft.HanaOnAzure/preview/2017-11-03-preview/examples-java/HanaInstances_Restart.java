/** Samples for HanaInstances Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_Restart.json
     */
    /**
     * Sample code: Restart a HANA instance.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void restartAHANAInstance(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.hanaInstances().restart("myResourceGroup", "myHanaInstance", com.azure.core.util.Context.NONE);
    }
}
