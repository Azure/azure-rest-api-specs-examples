/** Samples for Operations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Operations_Get.json
     */
    /**
     * Sample code: Operations_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void operationsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.operations().getWithResponse("{locationName}", "{operationName}", com.azure.core.util.Context.NONE);
    }
}
