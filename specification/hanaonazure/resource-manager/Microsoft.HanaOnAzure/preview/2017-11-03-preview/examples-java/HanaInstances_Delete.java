/** Samples for HanaInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaInstances_Delete.json
     */
    /**
     * Sample code: Delete a HANA instance.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void deleteAHANAInstance(com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.hanaInstances().delete("myResourceGroup", "myHanaInstance", com.azure.core.util.Context.NONE);
    }
}
