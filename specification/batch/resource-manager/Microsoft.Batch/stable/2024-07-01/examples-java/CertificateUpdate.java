
import com.azure.resourcemanager.batch.models.Certificate;

/**
 * Samples for Certificate Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/CertificateUpdate.json
     */
    /**
     * Sample code: UpdateCertificate.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void updateCertificate(com.azure.resourcemanager.batch.BatchManager manager) {
        Certificate resource = manager.certificates().getWithResponse("default-azurebatch-japaneast", "sampleacct",
            "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e", com.azure.core.util.Context.NONE).getValue();
        resource.update().withData("MIIJsgIBAzCCCW4GCSqGSIb3DQE...").withPassword("<ExamplePassword>").apply();
    }
}
