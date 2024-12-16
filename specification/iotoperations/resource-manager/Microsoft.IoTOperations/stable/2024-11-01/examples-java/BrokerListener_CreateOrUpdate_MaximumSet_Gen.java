
import com.azure.resourcemanager.iotoperations.models.BrokerListenerProperties;
import com.azure.resourcemanager.iotoperations.models.BrokerProtocolType;
import com.azure.resourcemanager.iotoperations.models.CertManagerCertificateSpec;
import com.azure.resourcemanager.iotoperations.models.CertManagerIssuerKind;
import com.azure.resourcemanager.iotoperations.models.CertManagerIssuerRef;
import com.azure.resourcemanager.iotoperations.models.CertManagerPrivateKey;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.ListenerPort;
import com.azure.resourcemanager.iotoperations.models.PrivateKeyAlgorithm;
import com.azure.resourcemanager.iotoperations.models.PrivateKeyRotationPolicy;
import com.azure.resourcemanager.iotoperations.models.SanForCert;
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
     * x-ms-original-file: 2024-11-01/BrokerListener_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerListener_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerListenerCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new BrokerListenerProperties().withServiceName("tpfiszlapdpxktx")
                .withPorts(Arrays.asList(new ListenerPort().withAuthenticationRef("tjvdroaqqy")
                    .withAuthorizationRef("fakeTokenPlaceholder").withNodePort(7281).withPort(1268)
                    .withProtocol(BrokerProtocolType.MQTT)
                    .withTls(new TlsCertMethod().withMode(TlsCertMethodMode.AUTOMATIC)
                        .withCertManagerCertificateSpec(new CertManagerCertificateSpec().withDuration("qmpeffoksron")
                            .withSecretName("fakeTokenPlaceholder").withRenewBefore("hutno")
                            .withIssuerRef(new CertManagerIssuerRef().withGroup("jtmuladdkpasfpoyvewekmiy")
                                .withKind(CertManagerIssuerKind.ISSUER).withName("ocwoqpgucvjrsuudtjhb"))
                            .withPrivateKey(new CertManagerPrivateKey().withAlgorithm(PrivateKeyAlgorithm.EC256)
                                .withRotationPolicy(PrivateKeyRotationPolicy.ALWAYS))
                            .withSan(new SanForCert().withDns(Arrays.asList("xhvmhrrhgfsapocjeebqtnzarlj"))
                                .withIp(Arrays.asList("zbgugfzcgsmegevzktsnibyuyp"))))
                        .withManual(new X509ManualCertificate().withSecretRef("fakeTokenPlaceholder")))))
                .withServiceType(ServiceType.CLUSTER_IP))
            .create();
    }
}
