import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void operationsList(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.operations().list(Context.NONE);
    }
}
