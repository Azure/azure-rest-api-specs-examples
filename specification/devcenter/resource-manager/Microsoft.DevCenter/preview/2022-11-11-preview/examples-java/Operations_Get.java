import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Operations_Get.json
     */
    /**
     * Sample code: Operations_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void operationsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.operations().list(Context.NONE);
    }
}
