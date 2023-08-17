/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void operations(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
