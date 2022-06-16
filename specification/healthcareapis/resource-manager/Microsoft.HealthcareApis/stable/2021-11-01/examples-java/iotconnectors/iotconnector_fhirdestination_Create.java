import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.healthcareapis.models.IotIdentityResolutionType;
import com.azure.resourcemanager.healthcareapis.models.IotMappingProperties;
import java.io.IOException;

/** Samples for IotConnectorFhirDestination CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
     */
    /**
     * Sample code: Create or update an Iot Connector FHIR destination.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateAnIotConnectorFHIRDestination(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) throws IOException {
        manager
            .iotConnectorFhirDestinations()
            .define("dest1")
            .withExistingIotconnector("testRG", "workspace1", "blue")
            .withResourceIdentityResolutionType(IotIdentityResolutionType.CREATE)
            .withFhirServiceResourceId(
                "subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice")
            .withFhirMapping(
                new IotMappingProperties()
                    .withContent(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"template\":[{\"template\":{\"codes\":[{\"code\":\"8867-4\",\"display\":\"Heart"
                                    + " rate\",\"system\":\"http://loinc.org\"}],\"periodInterval\":60,\"typeName\":\"heartrate\",\"value\":{\"defaultPeriod\":5000,\"unit\":\"count/min\",\"valueName\":\"hr\",\"valueType\":\"SampledData\"}},\"templateType\":\"CodeValueFhir\"}],\"templateType\":\"CollectionFhirTemplate\"}",
                                Object.class,
                                SerializerEncoding.JSON)))
            .withRegion("westus")
            .create();
    }
}
