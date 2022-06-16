import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to OepManager.
     */
    public static void operationsList(com.azure.resourcemanager.oep.OepManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
