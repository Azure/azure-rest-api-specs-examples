
import com.azure.resourcemanager.communication.models.VerificationParameter;
import com.azure.resourcemanager.communication.models.VerificationType;

/**
 * Samples for Domains CancelVerification.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/domains/
     * cancelVerification.json
     */
    /**
     * Sample code: Cancel verification.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void cancelVerification(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().cancelVerification("MyResourceGroup", "MyEmailServiceResource", "mydomain.com",
            new VerificationParameter().withVerificationType(VerificationType.SPF), com.azure.core.util.Context.NONE);
    }
}
