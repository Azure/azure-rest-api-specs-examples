/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseSkusList.json
     */
    /**
     * Sample code: SkusList.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void skusList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.skus().list("westus2", com.azure.core.util.Context.NONE);
    }
}
