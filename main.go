/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/letanthang/cli_search/cmd"
	organization_repository "github.com/letanthang/cli_search/organization/repository"
	organization_usecase "github.com/letanthang/cli_search/organization/usecase"
	"github.com/letanthang/cli_search/route"
	ticket_repository "github.com/letanthang/cli_search/ticket/repository"
	ticket_usecase "github.com/letanthang/cli_search/ticket/usecase"
	user_repository "github.com/letanthang/cli_search/user/repository"
	user_usecase "github.com/letanthang/cli_search/user/usecase"
)

func main() {
	organizationRepo := organization_repository.New()
	userRepo := user_repository.New()
	ticketRepo := ticket_repository.New()

	organizationUsecase := organization_usecase.New(organizationRepo, ticketRepo, userRepo)
	userUsecase := user_usecase.New(userRepo, ticketRepo, organizationRepo)
	ticketUsecase := ticket_usecase.New(ticketRepo, userRepo, organizationRepo)

	route.OrganizationUsecase = organizationUsecase
	route.TicketUsecase = ticketUsecase
	route.UserUsecase = userUsecase

	cmd.Execute()
}
