
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List_0.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void operationsList0(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
