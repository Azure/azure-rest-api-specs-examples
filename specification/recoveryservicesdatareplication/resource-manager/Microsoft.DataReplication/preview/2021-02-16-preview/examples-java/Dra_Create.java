import com.azure.resourcemanager.recoveryservicesdatareplication.models.DraModelCustomProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.DraModelProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.IdentityModel;

/** Samples for Dra Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Dra_Create.json
     */
    /**
     * Sample code: Dra_Create.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void draCreate(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .dras()
            .define("M")
            .withExistingReplicationFabric("rgrecoveryservicesdatareplication", "wPR")
            .withProperties(
                new DraModelProperties()
                    .withMachineId("envzcoijbqhtrpncbjbhk")
                    .withMachineName("y")
                    .withAuthenticationIdentity(
                        new IdentityModel()
                            .withTenantId("joclkkdovixwapephhxaqtefubhhmq")
                            .withApplicationId("cwktzrwajuvfyyymfstpey")
                            .withObjectId("khsiaqfbpuhp")
                            .withAudience("dkjobanyqgzenivyxhvavottpc")
                            .withAadAuthority("bubwwbowfhdmujrt"))
                    .withResourceAccessIdentity(
                        new IdentityModel()
                            .withTenantId("joclkkdovixwapephhxaqtefubhhmq")
                            .withApplicationId("cwktzrwajuvfyyymfstpey")
                            .withObjectId("khsiaqfbpuhp")
                            .withAudience("dkjobanyqgzenivyxhvavottpc")
                            .withAadAuthority("bubwwbowfhdmujrt"))
                    .withCustomProperties(new DraModelCustomProperties()))
            .create();
    }
}
