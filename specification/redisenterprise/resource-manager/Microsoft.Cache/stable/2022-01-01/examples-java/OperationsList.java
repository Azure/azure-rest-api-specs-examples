import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void operationsList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.operations().list(Context.NONE);
    }
}
