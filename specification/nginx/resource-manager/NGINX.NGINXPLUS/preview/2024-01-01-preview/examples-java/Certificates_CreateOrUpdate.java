
import com.azure.resourcemanager.nginx.models.NginxCertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/
     * Certificates_CreateOrUpdate.json
     */
    /**
     * Sample code: Certificates_CreateOrUpdate.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().define("default").withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .withProperties(new NginxCertificateProperties().withKeyVirtualPath("fakeTokenPlaceholder")
                .withCertificateVirtualPath("/src/cert/somePath.cert").withKeyVaultSecretId("fakeTokenPlaceholder"))
            .create();
    }
}
