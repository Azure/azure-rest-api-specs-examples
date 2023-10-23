/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Operations_Get.json
     */
    /**
     * Sample code: Operations_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void operationsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
