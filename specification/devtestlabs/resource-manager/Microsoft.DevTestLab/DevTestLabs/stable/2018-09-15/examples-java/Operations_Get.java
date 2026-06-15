
/**
 * Samples for Operations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Operations_Get.json
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
