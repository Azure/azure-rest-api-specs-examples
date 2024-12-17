
import com.azure.resourcemanager.iotoperations.models.BrokerListenerProperties;
import com.azure.resourcemanager.iotoperations.models.BrokerProtocolType;
import com.azure.resourcemanager.iotoperations.models.CertManagerCertificateSpec;
import com.azure.resourcemanager.iotoperations.models.CertManagerIssuerKind;
import com.azure.resourcemanager.iotoperations.models.CertManagerIssuerRef;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.ListenerPort;
import com.azure.resourcemanager.iotoperations.models.ServiceType;
import com.azure.resourcemanager.iotoperations.models.TlsCertMethod;
import com.azure.resourcemanager.iotoperations.models.TlsCertMethodMode;
import com.azure.resourcemanager.iotoperations.models.X509ManualCertificate;
import java.util.Arrays;

/**
 * Samples for BrokerListener CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerListener_CreateOrUpdate_Complex.json
     */
    /**
     * Sample code: BrokerListener_CreateOrUpdate_Complex.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerListenerCreateOrUpdateComplex(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new BrokerListenerProperties()
                .withPorts(Arrays.asList(
                    new ListenerPort().withAuthenticationRef("example-authentication").withPort(8080)
                        .withProtocol(BrokerProtocolType.WEB_SOCKETS),
                    new ListenerPort().withAuthenticationRef("example-authentication").withPort(8443)
                        .withProtocol(BrokerProtocolType.WEB_SOCKETS)
                        .withTls(new TlsCertMethod().withMode(TlsCertMethodMode.AUTOMATIC)
                            .withCertManagerCertificateSpec(new CertManagerCertificateSpec()
                                .withIssuerRef(new CertManagerIssuerRef().withGroup("jtmuladdkpasfpoyvewekmiy")
                                    .withKind(CertManagerIssuerKind.ISSUER).withName("example-issuer")))),
                    new ListenerPort().withAuthenticationRef("example-authentication").withPort(1883),
                    new ListenerPort().withAuthenticationRef("example-authentication").withPort(8883)
                        .withTls(new TlsCertMethod().withMode(TlsCertMethodMode.MANUAL)
                            .withManual(new X509ManualCertificate().withSecretRef("fakeTokenPlaceholder")))))
                .withServiceType(ServiceType.LOAD_BALANCER))
            .create();
    }
}
