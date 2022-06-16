import com.azure.core.util.Context;

/** Samples for OperationsStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/OperationsStatusGet.json
     */
    /**
     * Sample code: OperationsStatusGet.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void operationsStatusGet(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.operationsStatus().getWithResponse("West US", "testoperationid", Context.NONE);
    }
}
