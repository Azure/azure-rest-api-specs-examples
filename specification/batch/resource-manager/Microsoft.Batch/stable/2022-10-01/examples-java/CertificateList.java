import com.azure.core.util.Context;

/** Samples for Certificate ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/CertificateList.json
     */
    /**
     * Sample code: ListCertificates.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listCertificates(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .listByBatchAccount("default-azurebatch-japaneast", "sampleacct", 1, null, null, Context.NONE);
    }
}
