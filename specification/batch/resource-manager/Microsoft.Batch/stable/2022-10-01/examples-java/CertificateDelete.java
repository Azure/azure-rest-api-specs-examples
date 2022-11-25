import com.azure.core.util.Context;

/** Samples for Certificate Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/CertificateDelete.json
     */
    /**
     * Sample code: CertificateDelete.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void certificateDelete(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .delete(
                "default-azurebatch-japaneast",
                "sampleacct",
                "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e",
                Context.NONE);
    }
}
