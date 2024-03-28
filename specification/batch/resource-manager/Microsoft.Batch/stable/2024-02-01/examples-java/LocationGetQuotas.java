
/**
 * Samples for Location GetQuotas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/LocationGetQuotas.json
     */
    /**
     * Sample code: LocationGetQuotas.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void locationGetQuotas(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations().getQuotasWithResponse("japaneast", com.azure.core.util.Context.NONE);
    }
}
