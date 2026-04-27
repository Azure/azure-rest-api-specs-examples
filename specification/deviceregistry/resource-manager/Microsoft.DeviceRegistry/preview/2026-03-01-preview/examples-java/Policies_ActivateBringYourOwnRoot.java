
import com.azure.resourcemanager.deviceregistry.models.ActivateBringYourOwnRootRequest;

/**
 * Samples for Policies ActivateBringYourOwnRoot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Policies_ActivateBringYourOwnRoot.json
     */
    /**
     * Sample code: Policies_ActivateBringYourOwnRoot.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        policiesActivateBringYourOwnRoot(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().activateBringYourOwnRoot("rgdeviceregistry", "mynamespace", "mypolicy",
            new ActivateBringYourOwnRootRequest().withCertificateChain(
                "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAwIBAgIJAKL0UG+mRkmWMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjQwMTAxMDAwMDAwWhcNMjUwMTAxMDAwMDAwWjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAw...\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAwIBAgIJAKL0UG+mRkmXMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjQwMTAxMDAwMDAwWhcNMjUwMTAxMDAwMDAwWjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAx...\n-----END CERTIFICATE-----"),
            com.azure.core.util.Context.NONE);
    }
}
