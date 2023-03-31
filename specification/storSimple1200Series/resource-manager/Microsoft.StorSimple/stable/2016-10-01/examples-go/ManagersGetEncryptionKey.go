package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersGetEncryptionKey.json
func ExampleManagersClient_GetEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().GetEncryptionKey(ctx, "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SymmetricEncryptedSecret = armstorsimple1200series.SymmetricEncryptedSecret{
	// 	EncryptionAlgorithm: to.Ptr(armstorsimple1200series.EncryptionAlgorithmAES256),
	// 	Value: to.Ptr("EAAAAEVOuLjZQYRnJxD6RjOyBWGYfyw8wqXci5WHh8QfS98ouU65SFHHVWs2MdD2KeND3ZFM989gwfMaOVwiSmdMvi30/VnoNZIu7rhGG6SKVhUSHfOSZyb2eUufvwPXNUKa/mhVPJ4SH1uHUHvG1bjWaFZ2ojo1ff7e0xQlifYWZQdSi6ScxzjI41EdIT7Hspp+xtU2y+8Q5ALDgASRdVHdHYRmSF1/uyDXoAU8spTLrm/Ug8X0cadt1w+pAX0fnx3PPyfsNRVsbWofLtm1CirKv2euA8TamgFz82/xI4vT1m2RopLFc2W6sXGeSESWK2fUlV6WLjTqPwGXSJ7ZQ5/QcIP08QC1Z3K7muemFtSx+sr8/kjQIufxjD/A7cTN77bpTBCU7l1GxmdPFMlZxsVrU23SXAdu4JWcw6KrdSlxjig6GBHCHqCtaocjpD46GkRiGye80JudJqroWz0F14X9eOa48He3K/HLZnRdmaNKHClzApW3hCZwiI1r0NhHi9mLn3Laiy920L+kLRDghRsps9gMpYbQFYNNnQMAnCdWQ36TUbQWqqqcLVpgxaKvzbsXpYzL6ntlztmIp9RZ+i2r9ZoGGg4kIkiUNhpSVZ8k05H49zc1fJKapVER1MbEKApVBiC9ck+TbNMjkzOHY39QBJK35EYz2qfkGPTIDrnpdK4GiBXEDnq9ERx0CHVz8qLIRn538pZp8jktHft5WvOWk0Zw31+lSPnLCX1qBfnqT6ulH08VozwJqidcb2fOfMJ8BRgYupP47RGUa6gdRMMUmS3KgDetX0qXaqO4krCQYsiA3PeQwuZWVGA3eeALUQsizDFS1+esWVP/z9HgJclv4ydHjrMTf+GyWVozGJui9zqMz3B2otWZmWcyWAe4iE9a2w7P1MLgGvSShhQSU7S54mR4QF7WOArpaYUqe4VoZdLaTHS9bVDBrjPbFrrjfQlYNt8CPN2KmpeGsdicNlCeq5JCG9Ys7rpjr1QM49V0yV6sfOgE8rYldebwtJGwDnuq09LrGAIfuzVQ59AjW6YScPBSNeCnsyk4bC1I1zDKdEhCHspR8/3w++dAMH3wZ6fCEbOLgTy06lHVPyKyINVcDY/TLbSz04v9gFySujzSrt1qCrL5aK4eGxsANlVVwxXIGlTcyH5MYlrnNCIToIrdfIqmkjTea9WJ4buaz3DAnDsA+4ai+3vYMwJakE2mRxW5YIQHNEgoaY460a5HRCmnjDvpAmQ6ICB7uR/JunfQ2Jc3PI889dY8YOfy8YWHqzbXMUN7eRWRTp6A1W6J2CFtKLflDu0t+ZGX4Pa27zK4pN6ml410hoYGjet3+O7bxL/z3aR9cOgmV6kMhboOhvMaj8V43zSiq9ONhZgDZmVzhwP4oaKP/O1uhWqwbhGVOeAk8hCHTiD3FUXlrJVm3IgUSn8lL5R82cwyol578BKernVBP5PT4vwyF1aWPaOyxx+kITBcALBmlQv3JanW5j9FkFP0gC5oy7KitEXYac132oe7HiOQ7t1TJCoem5kxK3bYLg=="),
	// 	ValueCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
	// }
}
