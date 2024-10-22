
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticationMethod;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticationProperties;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodSat;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodX509;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethodX509Attributes;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthenticatorMethods;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BrokerAuthentication CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/BrokerAuthentication_CreateOrUpdate_Complex.json
     */
    /**
     * Sample code: BrokerAuthentication_CreateOrUpdate_Complex.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerAuthenticationCreateOrUpdateComplex(
        com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthentications().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new BrokerAuthenticationProperties()
                    .withAuthenticationMethods(
                        Arrays
                            .asList(
                                new BrokerAuthenticatorMethods()
                                    .withMethod(BrokerAuthenticationMethod.SERVICE_ACCOUNT_TOKEN)
                                    .withServiceAccountTokenSettings(new BrokerAuthenticatorMethodSat()
                                        .withAudiences(Arrays.asList("aio-internal"))),
                                new BrokerAuthenticatorMethods().withMethod(BrokerAuthenticationMethod.X509)
                                    .withX509Settings(new BrokerAuthenticatorMethodX509()
                                        .withAuthorizationAttributes(mapOf("root",
                                            new BrokerAuthenticatorMethodX509Attributes()
                                                .withAttributes(mapOf("organization", "contoso"))
                                                .withSubject("CN = Contoso Root CA Cert, OU = Engineering, C = US"),
                                            "intermediate",
                                            new BrokerAuthenticatorMethodX509Attributes()
                                                .withAttributes(mapOf("city", "seattle", "foo", "bar"))
                                                .withSubject("CN = Contoso Intermediate CA"),
                                            "smart-fan",
                                            new BrokerAuthenticatorMethodX509Attributes()
                                                .withAttributes(mapOf("building", "17")).withSubject("CN = smart-fan")))
                                        .withTrustedClientCaCert("my-ca")))))
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
