
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticationMethod;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticationProperties;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorCustomAuth;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodCustom;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodSat;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodX509;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodX509Attributes;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethods;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.X509ManualCertificate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BrokerAuthentication CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/BrokerAuthentication_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthentication_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthenticationCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthentications().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new BrokerAuthenticationProperties()
                    .withAuthenticationMethods(
                        Arrays
                            .asList(
                                new BrokerAuthenticatorMethods().withMethod(BrokerAuthenticationMethod.CUSTOM)
                                    .withCustomSettings(
                                        new BrokerAuthenticatorMethodCustom()
                                            .withAuth(new BrokerAuthenticatorCustomAuth().withX509(
                                                new X509ManualCertificate().withSecretRef("fakeTokenPlaceholder")))
                                            .withCaCertConfigMap("pdecudefqyolvncbus")
                                            .withEndpoint("https://www.example.com").withHeaders(mapOf("key8518",
                                                "fakeTokenPlaceholder")))
                                    .withServiceAccountTokenSettings(
                                        new BrokerAuthenticatorMethodSat().withAudiences(Arrays.asList("jqyhyqatuydg")))
                                    .withX509Settings(new BrokerAuthenticatorMethodX509()
                                        .withAuthorizationAttributes(mapOf("key3384",
                                            new BrokerAuthenticatorMethodX509Attributes()
                                                .withAttributes(mapOf("key186", "fakeTokenPlaceholder"))
                                                .withSubject("jpgwctfeixitptfgfnqhua")))
                                        .withTrustedClientCaCert("vlctsqddl")))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
