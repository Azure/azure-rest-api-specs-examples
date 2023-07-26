/** Samples for Certificate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/CertificateGet.json
     */
    /**
     * Sample code: Get Certificate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void getCertificate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .getWithResponse(
                "default-azurebatch-japaneast",
                "sampleacct",
                "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e",
                com.azure.core.util.Context.NONE);
    }
}
