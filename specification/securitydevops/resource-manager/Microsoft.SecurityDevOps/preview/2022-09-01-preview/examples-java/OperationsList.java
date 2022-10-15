import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/OperationsList.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void operationsList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
