
import com.azure.resourcemanager.sphere.models.ProofOfPossessionNonceRequest;

/**
 * Samples for Certificates RetrieveProofOfPossessionNonce.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/
     * PostRetrieveProofOfPossessionNonce.json
     */
    /**
     * Sample code: Certificates_RetrieveProofOfPossessionNonce.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void
        certificatesRetrieveProofOfPossessionNonce(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.certificates().retrieveProofOfPossessionNonceWithResponse("MyResourceGroup1", "MyCatalog1", "active",
            new ProofOfPossessionNonceRequest().withProofOfPossessionNonce("proofOfPossessionNonce"),
            com.azure.core.util.Context.NONE);
    }
}
